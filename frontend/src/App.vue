<template>
  <v-app id="movie_list">
    <v-app-bar :elevation="2">
      <template v-slot:prepend>
        <v-app-bar-nav-icon></v-app-bar-nav-icon>
      </template>
      <v-app-bar-title>Movie Night</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <Suspense>
        <MovieList />
      </Suspense>
    </v-main>
  </v-app>
</template>

<script>
import MovieList from './components/MovieList.vue'

export default {
  name: 'App',
  components: {
    MovieList
  }
}
</script>