<template>
  <q-page>
    <iframe id="map" frameborder="0" :src="mapSrc"></iframe>
    <div class="row justify-center">
      <q-card flat itemscope itemtype="http://schema.org/Event" class="col col-md-6">
        <q-card-section>
          <div itemprop="name" class="text-h5">{{item.name}}</div>
          <div>
            <q-icon name="event"/>
            <span itemprop="startDate">{{date.formatDate(item.startDate, 'YYYY-MM-DD')}}</span>({{date.formatDate(item.startDate, 'dddd')}})
            <span v-show="item.startDate !== item.endDate">
            - <span itemprop="endDate">{{date.formatDate(item.endDate, 'YYYY-MM-DD')}}</span>({{date.formatDate(item.endDate, 'dddd')}})
            </span>
          </div>
          <div itemscope itemprop="location" itemtype="http://schema.org/Place">
            <q-icon name="place"/>
            <span itemprop="address">{{item.location}}</span>
          </div>
          <div v-if="item.url"><q-icon name="home"/><a itemprop="url" :href="item.url">{{item.url}}</a></div>
        </q-card-section>
        <q-card-section>
          <div itemprop="description" v-html="item.detail"></div>
        </q-card-section>
      </q-card>
    </div>
  </q-page>
</template>

<script setup>
import { date } from 'quasar'
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { eventService } from 'src/lib/service'

const $route = useRoute()
const item = reactive({})
const mapSrc = ref('')

const getEvent = async (id) => {
  const response = await eventService.getEvent({ id })
  Object.assign(item, response.event)
  item.startDate = item.startDate.toDate()
  item.endDate = item.endDate.toDate()
  mapSrc.value = `https://www.google.com/maps/embed/v1/place?key=${process.env.GOOGLE_MAPS_API_KEY}&zoom=${item.zoom}&q=${item.lat},${item.lng}`
}

onMounted(() => {
  getEvent($route.params.id)
})
</script>

<style scoped>
#map {
  width: 100%;
  min-height: 25rem;
}
</style>
