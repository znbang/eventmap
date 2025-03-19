<template>
  <q-pagination v-model="page" :max="total" :to-fn="pageLink" input size="1.5rem" class="justify-center" />
</template>

<script setup>
import { ref } from 'vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'

const $route = useRoute()
const page = ref(parseInt($route.params.page || $route.query.page || 1))

defineProps({
  total: {
    type: Number,
    required: true,
  }
})

function pageLink(page) {
  return $route.params.page ? {params: {...$route.params, page}} : {query: {...$route.query, page}}
}

onBeforeRouteUpdate((to) => {
  page.value = parseInt(to.params.page || to.query.page || 1)
})
</script>
