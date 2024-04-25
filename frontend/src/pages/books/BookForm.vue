<template>
  <div class="row justify-center q-mt-md">
    <q-form class="col col-md-6 q-gutter-md" @submit="onSubmit">
      <q-input v-model="model.title" :label="$t('books.title')" :error="!!errors.title" :error-message="errors.title" outlined />
      <q-input v-model="model.author" :label="$t('books.author')" :error="!!errors.author" :error-message="errors.author" outlined />
      <q-input v-model="model.url" :label="$t('books.url')" :error="!!errors.url" :error-message="errors.url" outlined />
      <q-btn type="submit" color="primary" :label="$t('save')" />
      <q-btn flat to="/books" :label="$t('cancel')" />
    </q-form>
  </div>
</template>

<script setup>
import { tr } from 'boot/i18n'
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { useRoute, useRouter } from 'vue-router'
import { bookService } from 'src/lib/service'
import { Code, ConnectError } from '@connectrpc/connect'
import { ValidationError } from 'src/gen/errdetails/validation_pb'

const $q = useQuasar()
const $route = useRoute()
const $router = useRouter()
const errors = ref({})
const model = ref({})

async function fetchData() {
  const id = $route.params.id
  if (id) {
    const response = await bookService.getBook({ id })
    model.value = response.book
  }
}

async function onSubmit() {
  const id = $route.params.id
  const fn = id ? bookService.updateBook : bookService.createBook
  fn({ ...model.value })
    .then(() => $router.push('/books'))
    .catch(e => {
      if (e instanceof ConnectError && e.code === Code.InvalidArgument) {
        const err = e.findDetails(ValidationError).find(it => it.errors)
        errors.value = tr(err?.errors || {})
      } else {
        $q.notify(e)
      }
    })
}

fetchData()
</script>
