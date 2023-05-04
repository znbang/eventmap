<template>
  <q-pagination v-model="page" :max="total" @update:model-value="onUpdatePage" input size="1.5rem" class="justify-center" />
</template>

<script setup>
import { ref } from 'vue'
import { onBeforeRouteUpdate, useRoute, useRouter } from 'vue-router'

const $route = useRoute()
const $router = useRouter()
const page = ref(parseInt($route.params.page || $route.query.page || 1))
const props = defineProps({
  total: {
    type: Number,
    required: true,
  }
})

function onUpdatePage(page) {
  $router.push($route.params.page ? {params: {...$route.params, page}} : {query: {...$route.query, page}})
}

onBeforeRouteUpdate((to) => {
  page.value = parseInt(to.params.page || to.query.page || 1)
})
</script>
