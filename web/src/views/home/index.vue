
<template>
    <div class="home-container">
       <el-container>
            <el-aside width="200px" class="home-aside">
                <div>123</div>
            </el-aside>
            <el-container>
                <el-header>Header</el-header>
                <el-main>Main</el-main>
            </el-container>
        </el-container>
    </div>
</template>


<script setup lang="ts">

</script>

<style scoped lang="scss">

.home-container {
    width: 100%;
    height: 100%;
    min-height: 100vh;
}

.home-aside {
    width: 100%;
    height: 100vh;
}

</style>