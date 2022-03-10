<template>
  <q-page>
    <book-form />
  </q-page>
</template>

<script>
import BookForm from './BookForm'

export default {
  name: "NewBookPage",
  components: {
    BookForm
  }
}
</script>
