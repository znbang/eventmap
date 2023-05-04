<template>
  <q-page class="row justify-center" v-touch-swipe.left.right="swipePage">
    <q-list class="col col-md-8" separator>
      <q-item v-for="item in state.items" :key="item.id" :to="`/events/${item.id}`">
        <q-item-section>
          <q-item-label caption>{{ dateRange(item.startDate, item.endDate) }}</q-item-label>
          <q-item-label>{{ item.name }}</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
    <q-page-sticky position="bottom-right" :offset="state.fabOffset" v-touch-pan.prevent.mouse="dragFab">
      <q-btn fab icon="add" color="primary" to="/events/new" :disable="state.dragging" />
    </q-page-sticky>
  </q-page>
  <q-footer class="bg-white">
    <route-pagination :total="state.total"></route-pagination>
  </q-footer>
</template>

<script setup>
import { date } from 'quasar'
import { reactive } from 'vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
import { useSwipePage } from 'src/lib/swipe'
import { eventService } from 'src/lib/service'
import RoutePagination from 'components/RoutePagination.vue'

const $route = useRoute()
const state = reactive({
  fabOffset: [18, 18],
  dragging: false,
  items: [],
  total: 0,
})
const swipePage = useSwipePage(state)

function dragFab(evt) {
  state.dragging = evt.isFirst !== true && evt.isFinal !== true

  state.fabOffset = [
    state.fabOffset[0] - evt.delta.x,
    state.fabOffset[1] - evt.delta.y
  ]
}

async function updateState(to = $route) {
  const params = {
    page: parseInt(to.query.page || 1),
    filter: to.query.q || '',
  }
  const { items, total } = await eventService.listEvent(params)
  items.forEach(item => {
    item.startDate = item.startDate.toDate()
    item.endDate = item.endDate.toDate()
  })
  Object.assign(state, { items, total })
}

function dateRange(startDate, endDate) {
  if (startDate === endDate) {
    return date.formatDate(startDate, 'YYYY-MM-DD')
  } else {
    return date.formatDate(startDate, 'YYYY-MM-DD') + ' - ' + date.formatDate(endDate, 'YYYY-MM-DD')
  }
}

onBeforeRouteUpdate(updateState)
updateState()
</script>
