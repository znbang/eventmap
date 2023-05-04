<template>
  <q-page class="row justify-center" v-touch-swipe.left.right="swipePage">
    <q-list class="col col-md-8" separator>
      <div class="text-h6">
        <q-btn flat icon="arrow_back" to="/books" />
        {{ state.book.title }}
      </div>
      <hr>
      <q-item v-for="item in state.items" :key="item.id" :to="`/books/${state.book.id}/${item.page}`">
        <q-item-section>
          <q-item-label>{{ item.title }}</q-item-label>
        </q-item-section>
        <q-item-section side>
          <q-item-label>{{ item.page }}</q-item-label>
        </q-item-section>
        <q-item-section avatar side>
          <q-btn icon="delete" dense flat round @click.prevent="deletePage(item.page)"></q-btn>
        </q-item-section>
      </q-item>
    </q-list>
  </q-page>
  <q-footer class="bg-white">
    <route-pagination :total="state.total"></route-pagination>
  </q-footer>
</template>

<script setup>
import { useQuasar } from "quasar";
import { reactive } from 'vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
import { useSwipePage } from 'src/lib/swipe'
import { bookService } from 'src/lib/service'
import { useI18n } from "vue-i18n";
import RoutePagination from 'components/RoutePagination.vue'

const $q = useQuasar()
const $route = useRoute()
const { t } = useI18n()
const state = reactive({
  book: {
    id: '',
  },
  items: [],
  total: 0,
})
const swipePage = useSwipePage(state)

async function updateState(to = $route) {
  const params = {
    id: to.params.id,
    page: parseInt(to.query.page || 1)
  }
  const { book, total, items } = await bookService.getToc(params)
  Object.assign(state, { book, items, total })
}

function deletePage(page) {
  $q.dialog({
    message: t('chapters.dialog.confirm'),
    cancel: { flat: true, label: t('cancel') },
    ok: { color: 'negative', label: t('delete') }
  }).onOk(async () => {
    await bookService.deleteChapter({ id: state.book.id, page })
    await updateState()
    $q.notify(t('chapters.dialog.deleted'))
  })
}

onBeforeRouteUpdate(updateState)
updateState()
</script>
