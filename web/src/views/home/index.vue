
<template>
    <div class="home-container">
        <el-button @click="send">发送请求</el-button>
    </div>
</template>


<script setup lang="ts">
import { role } from "@/api/role"
    const send = () => {
        role()
    }
</script>