
<template>
  <div class="home-container">
    <el-container>
      <el-aside :width="main.coll ? '65px' : '200px' " class="home-aside">
        <div class="home-aside-head">1232</div>
        <aside-menu
          :menus="menuList"
          :text-color="main.textColor"
          :background-color="main.backgroundColor"
          :active-text-color="main.activeTextColor"
          :collapse="main.coll"
          :router="true"
        />
      </el-aside>
      <el-container>
        <el-header class="home-header">
          <div @click="main.toggleCollapse" class="head-btn">
            <el-icon v-if="main.coll" :size="iconSize">
              <expand />
            </el-icon>
            <el-icon v-else :size="iconSize">
              <fold />
            </el-icon>
          </div>
          <div class="home-head-right">
              <el-dropdown trigger="hover">
                <el-avatar :size="30" src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item>Action 1</el-dropdown-item>
                    <el-dropdown-item>Action 2</el-dropdown-item>
                    <el-dropdown-item>Action 3</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
          </div>
        </el-header>
        <el-main class="home-main" v-loading="main.loading">
          <el-scrollbar>
            <router-view />
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>


<script setup lang="ts">
import { Fold } from "@element-plus/icons-vue";
import { mainStore } from "@/pinia/index";
import AsideMenu from "@/components/asideMenus/index.vue";

const iconSize = 30;

const main = mainStore();

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
    name: "abc5",
    path: "/abc546",
    icon: "Search",
    hidden: false,
    label: "其他2",
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
  border-left: 1px solid var(--el-border-color);
}

.hcontainer {
  height: 100vh;
}

.head-btn {
  cursor: pointer;
  display: flex;

  align-items: center;
  justify-content: center;
}
</style>