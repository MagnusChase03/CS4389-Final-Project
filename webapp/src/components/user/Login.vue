<script setup>
import { ref } from 'vue';

let username = ref("");
let password = ref("");

function login() {
    fetch(import.meta.env.VITE_API_URL + "/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `username=${username.value}&password=${password.value}`,
        credentials: "include"
    }).then((resp) => {
        username.value = "";
        password.value = "";
    });
}
</script>

<template>
<div class="container-fluid">
<div class="row">
<div class="col-sm"></div> 
<div class="col-sm border border-dark p-5 rounded">
    <form>
        <input type="text" class="form-control" id="username" placeholder="Username" v-model="username">
        <input type="password" class="form-control" id="password" placeholder="Password" v-model="password">
        <div class="text-center">
            <button type="submit" class="btn btn-primary" @click.prevent="login">Login</button>
        </div>
    </form>
</div> 
<div class="col-sm"></div> 
</div>
</div>
</template>

<style scoped>
input {
    margin-bottom: 1rem;
}
</style>
