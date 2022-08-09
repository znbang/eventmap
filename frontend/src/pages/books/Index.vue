<template>
  <q-page class="row justify-center" v-touch-swipe.left.right="swipePage">
    <q-list class="col col-md-8" separator>
      <q-expansion-item group="toc" :content-inset-level="1" :icon="getIcon(item)" :class="{spinner: item.job?.status === 2}" v-for="item in state.items" :label="item.title" :caption="item.author" :to="`/books/${item.id}/toc`">
        <q-card>
          <q-card-section class="row items-center">
            <span class="text-grey-6">{{item.job?.message}}</span>
            <q-space />
            <q-btn dense flat round icon="edit" :to="`/books/${item.id}/edit`"><q-tooltip>{{ $t('edit') }}</q-tooltip></q-btn>
            <q-btn dense flat round icon="delete" @click="deleteBook(item.id)"><q-tooltip>{{ $t('delete') }}</q-tooltip></q-btn>
            <q-btn v-if="(item.job?.status || 0) === 0" dense flat round icon="refresh" @click="syncAll(item.id)"><q-tooltip>{{ $t('books.syncAll') }}</q-tooltip></q-btn>
            <q-btn v-if="(item.job?.status || 0) === 0" dense flat round icon="sync" @click="syncNew(item.id)"><q-tooltip>{{ $t('books.syncNew') }}</q-tooltip></q-btn>
            <q-btn v-if="(item.job?.status || 0) !== 0" dense flat round icon="stop_circle" @click="stopSync(item.id)" />
            <q-btn dense flat round icon="save_alt" @click="download(item.id)"><q-tooltip>{{ $t('books.download') }}</q-tooltip></q-btn>
          </q-card-section>
        </q-card>
      </q-expansion-item>
    </q-list>
    <q-page-sticky position="bottom-right" :offset="state.fabOffset" v-touch-pan.prevent.mouse="dragFab">
      <q-btn fab icon="add" color="primary" to="/books/new" :disable="state.dragging" />
    </q-page-sticky>
  </q-page>
  <q-footer class="bg-white">
    <z-pagination :total="state.total" />
  </q-footer>
</template>

<script setup>
import { useQuasar } from 'quasar'
import { reactive, watch, onMounted } from 'vue'
import { onBeforeRouteLeave, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useSwipePage } from 'components/swipe'
import { bookService } from 'components/service'
import ZPagination from 'components/ZPagination.vue'

const $q = useQuasar()
const $route = useRoute()
const { t } = useI18n()
const state = reactive({
  fabOffset: [18, 18],
  dragging: false,
  items: [],
  total: 0,
})
const swipePage = useSwipePage(state)
const abortController = new AbortController()
const signal = abortController.signal

function dragFab(evt) {
  state.dragging = evt.isFirst !== true && evt.isFinal !== true

  state.fabOffset = [
    state.fabOffset[0] - evt.delta.x,
    state.fabOffset[1] - evt.delta.y
  ]
}

async function updateState() {
  const params = {
    page: parseInt($route.query.page || 1),
    filter: $route.query.q || '',
  }
  const { items, total } = await bookService.listBook(params)
  Object.assign(state, { items, total })
}

async function syncAll(id) {
  await bookService.syncAll({ id })
  await updateState()
}

async function syncNew(id) {
  await bookService.syncNew({ id })
  await updateState()
}

async function stopSync(id) {
  await bookService.stopSync({ id })
  await updateState()
}

function deleteBook(id) {
  $q.dialog({
    message: t('books.dialog.confirm'),
    cancel: { flat: true, label: t('cancel') },
    ok: { color: 'negative', label: t('delete') }
  }).onOk(async () => {
    await bookService.deleteBook({ id })
    await updateState()
    $q.notify(t('books.dialog.deleted'))
  })
}

async function download(id) {
  const { title, content } = await bookService.downloadBook({ id })
  const blob = new Blob([content], { type: 'text/plain' })
  const a = document.createElement('a')
  a.href = URL.createObjectURL(blob)
  a.download = title + '.txt'
  a.click()
}

function getIcon(book) {
  switch (book.job?.status || 0) {
    case 1: return 'schedule'
    case 2: return 'sync'
    default: return 'done'
  }
}

function syncStatusStream() {
  (async() => {
    try {
      for await (const event of bookService.syncStatus({}, {signal})) {
        state.items.forEach(item => {
          if (item.job?.id === event.id) {
            item.job.message = event.message
            item.job.status = event.status
          }
        })
      }
    } catch (e) {}
  })()
}

function syncStatusSse() {
  const eventSource = new EventSource('/api/sync-book-status')
  eventSource.addEventListener('message', async (event) => {
    const task = JSON.parse(event.data)
    state.items.forEach(item => {
      if (item.job?.id === task.id) {
        item.job.message = task.message
        item.job.status = task.status
      }
    })
  })
}

// onMounted(syncStatusStream)
onMounted(syncStatusSse)

onBeforeRouteLeave(watch(() => $route.query, updateState))
onBeforeRouteLeave(() => abortController.abort())

updateState()
</script>

<style>
.spinner .q-item__section--avatar .q-icon {
  animation: rotating 2s infinite linear;
}
@keyframes rotating {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(-1turn);
  }
}
</style>
