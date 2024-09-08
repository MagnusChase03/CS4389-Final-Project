<script setup>
import { ref } from 'vue';

let status = ref("");
fetch(import.meta.env.VITE_API_URL + "/healthcheck").then((resp) => resp.json()).then((json) => {
    status.value = json.Data;
});
</script>

<template>
    <h1>Status: {{ status }}</h1>
</template>

<style scoped>
h1 {
    color: orange;
}
</style>
