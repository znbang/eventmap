import { useRoute, useRouter, onBeforeRouteLeave } from 'vue-router'

export function useSwipePage(state) {
  const $route = useRoute()
  const $router = useRouter()

  function paging(direction) {
    const params = !!$route.params.page
    let page = params ? parseInt($route.params.page || 1) : parseInt($route.query.page || 1)
    if (direction === 'left') {
      if (page < state.total) {
        page += 1
        $router.push(params ? {params: {...$route.params, page}} : {query: {...$route.query, page}})
      }
    } else if (direction === 'right') {
      if (page > 1) {
        page -= 1
        $router.push(params ? {params: {...$route.params, page}} : {query: {...$route.query, page}})
      }
    }
  }

  function onKeyUp(e) {
    const params = !!$route.params.page
    if (e.ctrlKey) {
      switch (e.key) {
        case 'Home': $router.push(params ? {params: {...$route.params, page: 1}} : {query: {...$route.query, page: 1}}); break;
        case 'End': $router.push(params ? {params: {...$route.params, page: state.total}} : {query: {...$route.query, page: state.total}}); break;
      }
    } else {
      switch (e.key) {
        case 'ArrowLeft': paging('right'); break;
        case 'ArrowRight': paging('left'); break;
      }
    }
  }

  document.addEventListener('keyup', onKeyUp)

  onBeforeRouteLeave(() => document.removeEventListener('keyup', onKeyUp))

  return ({ direction }) => {
    paging(direction)
  }
}
