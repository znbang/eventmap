<template>
  <q-page padding class="column items-center">
    <div class="column q-gutter-md">
      <q-btn type="a" color="primary" no-caps v-for="provider in state.items" :key="provider.id" :href="provider.url">{{ $t(`auths.${provider.name}`) }}</q-btn>
    </div>
  </q-page>
</template>

<script setup>
import { reactive } from 'vue'
import { authService } from 'components/service'

const state = reactive({
  items: {},
})

async function updateState() {
  const response = await authService.listProvider({})
  state.items = response.providers
}

updateState()
</script>
