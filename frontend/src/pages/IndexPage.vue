<template>
  <q-page class="row">
    <div id="map" class="col"></div>
  </q-page>
</template>

<script>
import { date, useQuasar } from 'quasar'
import { defineComponent, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Loader } from '@googlemaps/js-api-loader'
import { MarkerClusterer } from '@googlemaps/markerclusterer'
import { eventService } from 'src/lib/service'
import { timestampDate } from '@bufbuild/protobuf/wkt'

export default defineComponent({
  setup() {
    const $q = useQuasar()
    const $router = useRouter()
    const loader = new Loader({
      apiKey: process.env.GOOGLE_MAPS_API_KEY,
      libraries: ['places'],
    })

    function addMarkers(map, events, google) {
      const infoWindow = new google.maps.InfoWindow()
      const markers = events.map(event => {
        const marker = new google.maps.Marker({
          position: event,
          map: map,
          icon: { url: '/images/map-marker-icon.png', scaledSize: { width: 36, height: 36 } }
        })
        marker.addListener('mouseover', function () {
          let content = '<div>' + event.name + '</div><div><i class="material-icons">date_range</i>' + event.startDate
          if (event.startDate !== event.endDate) {
            content += ' ~ ' + event.endDate
          }
          content += '</div>'
          infoWindow.setContent(content)
          infoWindow.open(map, marker)
        })
        marker.addListener('mouseout', () => {
          infoWindow.close()
        })
        marker.addListener('click', () => {
          $router.push('/events/' + event.id)
        })
        return marker
      })

      new MarkerClusterer(map, markers, { imagePath: '/images/m' })
    }

    onMounted(async () => {
      loader
        .load()
        .then(async google => {
          const options = {
            zoom: 8,
            center: { lat: 23.7, lng: 121.2 },
            fullscreenControl: false,
            streetViewControl: false,
            mapTypeControlOptions: {
              position: google.maps.ControlPosition.RIGHT_TOP
            }
          }
          const map = new google.maps.Map(document.getElementById('map'), options)
          const response = await eventService.listActiveEvent({})
          const items = response.items || []
          const events = items.map(a => {
            a.startDate = date.formatDate(timestampDate(a.startDate), 'YYYY-MM-DD')
            a.endDate = date.formatDate(timestampDate(a.endDate), 'YYYY-MM-DD')
            return a
          })
          addMarkers(map, events, google)
        })
        .catch(error => {
          $q.notify({
            type: 'negative',
            message: error,
          })
        })
    })
  }
})
</script>
