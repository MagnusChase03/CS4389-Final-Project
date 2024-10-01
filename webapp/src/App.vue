<script setup>
import Navbar from './components/Navbar.vue';
import LoginPage from './components/user/Login.vue';

import { ref } from 'vue';

let currentPage = ref("login");
</script>

<template>
<Navbar/>
<LoginPage v-if="currentPage == 'login'" />
</template>

<style scoped>
</style>
