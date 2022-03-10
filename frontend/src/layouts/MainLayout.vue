<template>
  <q-layout view="lHh Lpr lFf">
    <q-header bordered class="bg-white text-primary">
      <q-toolbar class="row">
        <q-btn class="q-mr-sm" flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer"/>
        <q-btn class="gt-xs" flat dense no-caps stretch type="a" to="/">
          <q-toolbar-title>
            EventMap
          </q-toolbar-title>
        </q-btn>
        <q-space />
        <q-input class="col-xs col-sm-6 col-md-3 col-lg-2" v-model="state.q" clearable dense outlined :placeholder="$t('search')" @clear="$router.push({ query: {}})" @keyup.enter="$router.push({ query: { q: state.q }})">
          <template v-slot:after>
            <q-btn icon="search" dense flat round :to="{ query: { q: state.q } }"/>
          </template>
        </q-input>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      bordered
      show-if-above
      class="bg-grey-1"
    >
      <q-list>
        <q-item to="/" exact>
          <q-item-section avatar>
            <q-icon name="home" />
          </q-item-section>
          <q-item-section>{{$t('home')}}</q-item-section>
        </q-item>
        <q-item to="/events">
          <q-item-section avatar>
            <q-icon name="school" />
          </q-item-section>
          <q-item-section>{{$t('event')}}</q-item-section>
        </q-item>
        <q-item to="/events/user" v-if="isAuthenticated()">
          <q-item-section avatar>
            <q-icon name="library_add" />
          </q-item-section>
          <q-item-section>{{$t('manage')}}</q-item-section>
        </q-item>
        <q-item to="/books" v-if="isAuthenticated()">
          <q-item-section avatar>
            <q-icon name="menu_book" />
          </q-item-section>
          <q-item-section>{{$t('book')}}</q-item-section>
        </q-item>
        <q-separator />
        <q-expansion-item :label="$t('language')" icon="language">
          <q-list>
            <q-item clickable :inset-level="1" @click="$i18n.locale = 'zh-TW'">
              <q-item-section>中文</q-item-section>
            </q-item>
            <q-item clickable :inset-level="1" @click="$i18n.locale = 'en-US'">
              <q-item-section>English</q-item-section>
            </q-item>
          </q-list>
        </q-expansion-item>
        <q-separator />
        <q-item to="/login" v-if="!isAuthenticated()">
          <q-item-section avatar>
            <q-icon name="login" />
          </q-item-section>
          <q-item-section>{{$t('login')}}</q-item-section>
        </q-item>
        <q-item clickable @click="onClickLogout" v-if="isAuthenticated()">
          <q-item-section avatar>
            <q-icon name="logout" />
          </q-item-section>
          <q-item-section>{{$t('logout')}}</q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { useQuasar } from 'quasar'
import { isAuthenticated, logout } from 'boot/auth'
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

const leftDrawerOpen = ref(false)
const $router = useRouter()
const $route = useRoute()
const $q = useQuasar()
const { t } = useI18n()
const state = reactive({
  q: $route.query.q || '',
})

function onClickLogout() {
  logout().then(() => {
    $q.notify(t('登出成功。'))
    $router.push('/')
  })
}

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>
