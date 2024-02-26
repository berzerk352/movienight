<script setup>
import { ref } from 'vue'

const API_URL = 'http://localhost:3000'

const movies = ref(null)

movies.value = await (await fetch(`${API_URL}/movies`)).json()

const headers = [
  { key: 'id', title: 'ID' },
  { key: 'submitter', title: 'Submitter'},
  { key: 'movie_name', title: 'Movie'}
]
</script>

<template>
  <v-card flat title="Movies">
    <v-data-table :items="movies" :headers="headers"></v-data-table>
  </v-card>
</template>
  
<script>
export default {
  name: 'MovieList',
  // Fetch items from API or data store
}
</script>