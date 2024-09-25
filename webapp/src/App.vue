<script setup>
import Navbar from './components/Navbar.vue';
import LoginPage from './components/user/Login.vue';

let currentPage = "login";
</script>

<template>
<Navbar/>
<LoginPage v-if="currentPage" />
</template>

<style scoped>
</style>
