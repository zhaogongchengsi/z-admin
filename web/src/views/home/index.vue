
<template>
  <div class="home-container">
    <el-container>
      <el-aside :width="main.coll ? '65px' : '200px' " class="home-aside">
        <div class="home-aside-head">1232</div>
        <aside-menu
          :active-text-color="main.activeTextColor"
          :background-color="main.backgroundColor"
          :text-color="main.textColor"
          :collapse-transition="false"
          :collapse="main.coll"
          :menu-options="menuList"
        />
      </el-aside>
      <el-container>
        <el-header class="home-header">
          <el-button type="primary" :icon="Fold" @click="main.toggleCollapse" />
          <div class="home-head-right">123213</div>
        </el-header>
        <el-main class="home-main" v-loading="main.loading">
          <el-scrollbar>
            <div class="hcontainer">123213</div>
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>


<script setup lang="ts">
import { Fold } from "@element-plus/icons-vue";
import { mainStore } from "@/pinia/index";
import AsideMenu from "@/components/AsideMenu";

const menuList = [
  {
    name: "home",
    path: "/home",
    icon: "Search",
    label: "首页",
    hidden: true,
    children: [
      {
        name: "home2",
        path: "/home3",
        icon: "Search",
        label: "首页2",
        hidden: true,
      },
    ],
  },
  {
    name: "abc",
    path: "/abc",
    icon: "Search",
    label: "其他",
    hidden: true,
    children: [
      {
        name: "abc2",
        path: "/abc3",
        hidden: true,
        icon: "Search",
        label: "其他2",
      },
      {
        name: "abc5",
        path: "/abc5",
        icon: "Search",
        hidden: false,
        label: "其他2",
      },
    ],
  },
];

const main = mainStore();
</script>

<style scoped lang="scss">
.home-container {
  width: 100%;
  height: 100%;
  min-height: 100vh;
}

.home-aside {
  height: 100vh;
  background-color: v-bind("main.backgroundColor");
  // transition: all 0.3s;

  .home-aside-head {
    height: 60px;
  }
}

.home-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.home-main {
  --home-padding: 10px;
  height: calc(100vh - 60px);
  padding: var(--home-padding) 0 var(--home-padding) var(--home-padding);
  border-top: 1px solid var(--el-border-color);
}

.hcontainer {
  height: 100vh;
}
</style>