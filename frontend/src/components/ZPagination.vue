<template>
  <q-pagination v-model="page" :max="total" @update:model-value="onUpdatePage" input size="1.5rem" class="justify-center" />
</template>

<script setup>
import { ref, watch } from 'vue'
import { onBeforeRouteLeave, useRoute, useRouter } from 'vue-router'

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

onBeforeRouteLeave(watch(() => $route.query, () => {
  page.value = parseInt($route.params.page || $route.query.page || 1)
}))
</script>
