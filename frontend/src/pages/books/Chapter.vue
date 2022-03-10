<template>
  <q-page class="row justify-center" v-touch-swipe.left.right="swipePage">
    <div class="col col-md-6 q-mx-md text-large">
      <q-btn flat :to="`/books/${state.book.id}/toc`" style="float: right">{{ state.book.title }}</q-btn>
      <div class="text-h6">{{ state.chapter.title }}</div>
      <hr>
      <p v-for="line in state.chapter.body.split('\n')">{{ line }}</p>
    </div>
  </q-page>
  <q-footer class="bg-white">
    <z-pagination :total="state.total"></z-pagination>
  </q-footer>
</template>

<script setup>
import { reactive, watch } from 'vue'
import { onBeforeRouteLeave, useRoute } from 'vue-router'
import { useSwipePage } from 'components/swipe'
import { bookService } from 'components/service'
import ZPagination from 'components/ZPagination.vue'

const $route = useRoute()
const state = reactive({
  book: {
    id: $route.params.id,
  },
  chapter: {
    body: '',
  },
  total: 0,
})
const swipePage = useSwipePage(state)

async function updateState() {
  const page = parseInt($route.params.page || 1)
  const { book, chapter, total } = await bookService.getChapter({ id: state.book.id, page })
  Object.assign(state, { book, chapter, total })
}

onBeforeRouteLeave(watch(() => $route.params, updateState))

updateState()
</script>

<style scoped>
.text-large {
  font-size: large;
}
</style>
