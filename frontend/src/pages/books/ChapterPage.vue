<template>
  <q-page class="row justify-center" v-touch-swipe.left.right="swipePage">
    <div class="col col-md-6 q-mx-md text-large">
      <q-btn flat :to="`/books/${state.book.id}/toc`" style="float: right">{{ state.book.title }}</q-btn>
      <div class="text-h6">{{ state.chapter.title }}</div>
      <hr>
      <p v-for="(line, index) in state.chapter.body.split('\n')" :key="index">{{ line }}</p>
    </div>
  </q-page>
  <q-footer class="bg-white">
    <route-pagination :total="state.total"></route-pagination>
  </q-footer>
</template>

<script setup>
import { reactive } from 'vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
import { useSwipePage } from 'src/lib/swipe'
import { bookService } from 'src/lib/service'
import RoutePagination from 'components/RoutePagination.vue'

const $route = useRoute()
const state = reactive({
  book: {
    id: '',
  },
  chapter: {
    body: '',
  },
  total: 0,
})
const swipePage = useSwipePage(state)

async function updateState(to = $route) {
  const page = parseInt(to.params.page || 1)
  const { book, chapter, total } = await bookService.getChapter({ id: to.params.id, page })
  Object.assign(state, { book, chapter, total })
}

onBeforeRouteUpdate(updateState)
updateState()
</script>

<style scoped>
.text-large {
  font-size: large;
}
</style>
