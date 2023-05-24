<template>
  <div id="map"></div>
  <div class="row justify-center">
    <q-form class="col col-md-6 q-mt-xs q-mb-lg q-gutter-md" @submit="onSubmit()">
      <q-input :label="$t('events.location')" v-model="form.location" :error="!!errors.location" :error-message="errors.location" outlined />
      <q-input :label="$t('events.name')" v-model="form.name" :error="!!errors.name" :error-message="errors.name" outlined />
      <q-input :label="$t('events.date')" v-model="dateInput" :error="!!errors.startDate" :error-message="errors.startDate" outlined readonly>
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy ref="dateProxy">
              <q-date today-btn range mask="YYYY-MM-DD" v-model="dateRange" @update:model-value="onDateUpdate">
                <div class="row items-center justify-end">
                  <q-btn v-close-popup color="primary" flat :label="$t('close')" />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
      <q-input :label="$t('events.url')" v-model="form.url" :error="!!errors.url" :error-message="errors.url" outlined />
      <q-field :error="!!errors.detail" :error-message="errors.detail" borderless no-error-icon hide-bottom-space>
        <template v-slot:control>
          <div class="fit">
            <q-editor :placeholder="$t('events.detail')" v-model="form.detail" min-height="10rem" :toolbar="[['bold','italic','underline','strike','removeFormat'],['unordered', 'ordered'],['viewsource']]"/>
          </div>
        </template>
      </q-field>
      <q-btn type="submit" color="primary" :label="$t('save')"/>
      <q-btn to="/events/user" color="primary" flat class="q-ml-sm" :label="$t('cancel')"/>
    </q-form>
  </div>
  <div id="searchBoxWrapper" class="row justify-center">
    <input type="text" id="searchBox" class="q-field q-mb-md q-mx-md col col-md-6" :placeholder="$t('events.find')"/>
  </div>
</template>

<script setup>
import { date } from 'quasar'
import { tr } from 'boot/i18n'
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Loader } from '@googlemaps/js-api-loader'
import { eventService } from 'src/lib/service'
import { Timestamp } from '@bufbuild/protobuf'
import { ConnectError } from '@bufbuild/connect'
import { ValidationError } from 'src/gen/errdetails/validation_pb'

const $route = useRoute()
const $router = useRouter()
const dateInput = ref('')
const dateRange = ref(null)
const errors = ref({})
const form = reactive({
  detail: '',
  lat: 23.5,
  lng: 121.2,
  zoom: 8,
})
const loader = new Loader({
  apiKey: process.env.GOOGLE_MAPS_API_KEY,
  libraries: ['places'],
})

const initMap = (google) => {
  const options = {
    zoom: form.zoom,
    center: {lat: form.lat, lng: form.lng},
    gestureHandling: 'greedy',
    fullscreenControl: false,
    streetViewControl: false,
    mapTypeControlOptions: {position: google.maps.ControlPosition.RIGHT_TOP}
  }
  const map = new google.maps.Map(document.getElementById('map'), options)
  map.addListener('center_changed', () => {
    form.lat = map.getCenter().lat()
    form.lng = map.getCenter().lng()
  })
  map.addListener('zoom_changed', () => {
    form.zoom = map.getZoom()
  })
  // search box
  const searchBoxWrapper = document.getElementById('searchBoxWrapper')
  const searchBoxInput = document.getElementById('searchBox')
  const searchBox = new google.maps.places.SearchBox(searchBoxInput)
  map.controls[google.maps.ControlPosition.BOTTOM_CENTER].push(searchBoxWrapper)
  searchBox.addListener('places_changed', () => {
    const places = searchBox.getPlaces()
    if (places === null) {
      return
    }
    places.forEach(place => {
      if (!place.geometry) {
        return
      }
      map.setZoom(16)
      map.setCenter(place.geometry.location)
      form.location = searchBoxInput.value
    })
  })
}

const getEvent = async (id) => {
  const response = await eventService.getEvent({ id })
  Object.assign(form, response.event)
  const startDate = date.formatDate(form.startDate.toDate(), 'YYYY-MM-DD')
  const endDate = date.formatDate(form.endDate.toDate(), 'YYYY-MM-DD')
  if (startDate === endDate) {
    dateInput.value = startDate
    dateRange.value = startDate
  } else {
    dateInput.value = startDate + ' - ' + endDate
    dateRange.value = {
      from: startDate,
      to: endDate,
    }
  }
}

onMounted(() => {
  loader.load().then(google => {
    if ($route.params.id) {
      getEvent($route.params.id).then(() => initMap(google))
    } else {
      initMap(google)
    }
  })
})

function onDateUpdate(value) {
  if (value) {
    if (typeof value === 'string') {
      dateInput.value = value
      form.startDate = form.endDate = Timestamp.fromDate(new Date(date.formatDate(value, 'YYYY-MM-DDTHH:mm:ss.SSSZ')))
    } else {
      dateInput.value = value.from + ' - ' + value.to
      form.startDate = Timestamp.fromDate(new Date(date.formatDate(value.from, 'YYYY-MM-DDTHH:mm:ss.SSSZ')))
      form.endDate = Timestamp.fromDate(new Date(date.formatDate(value.to, 'YYYY-MM-DDTHH:mm:ss.SSSZ')))
    }
  } else {
    dateInput.value = ''
    form.startDate = Timestamp.fromDate(new Date(0))
    form.endDate = Timestamp.fromDate(new Date(0))
  }
}

function onSubmit() {
  const id = $route.params.id
  const fn = id ? eventService.updateEvent : eventService.createEvent
  fn({ ...form })
    .then(() => $router.push('/events/user'))
    .catch(e => {
      if (e instanceof ConnectError) {
        const err = e.findDetails(ValidationError).find(it => it.errors)
        errors.value = tr(err?.errors || {})
      }
    })
}
</script>

<style scoped>
#map {
  width: 100%;
  height: 25rem;
}
#map:after {
  position: absolute; top: 50%; left: 50%; width: 48px; height: 48px;
  margin: -48px 0 0 -24px; display: block; content: " ";
  background: url('/images/map-marker-icon.png');
  background-size: 48px 48px;
  pointer-events: none;
}
.q-field--error .q-editor {
  border: 1px solid red;
}
#searchBoxWrapper {
  width: 80%;
}
#searchBox {
  padding: 0.5rem;
}
</style>
