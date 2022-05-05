<template>
  <el-menu
    :text-color="porps.textColor"
    :collapse="porps.collapse"
    :background-color="porps.backgroundColor"
    :active-text-color="porps.activeTextColor"
    :router="porps.router"
    :collapse-transition="false"
    class="aside-menu"
  >
    <component
      v-for="(item) in porps.menus"
      :key="item.path"
      :is="iscomponent(item)"
      :index="item.path"
      :menuInfo="item"
    ></component>
  </el-menu>
</template>

<script lang="ts" setup>
import AsideItemMenu from "./itemMenu.vue";
import AsideSubMenu from "./subMenu.vue";

const porps = defineProps({
  menus: {
    type: Array,
    default: () => [],
  },
  textColor: {
    type: String,
    default: "",
  },
  backgroundColor: {
    type: String,
    default: "",
  },
  activeTextColor: {
    type: String,
    default: "",
  },
  collapse: {
    type: Boolean,
    default: false,
  },
  router: {
    type: Boolean,
    default: false,
  },
});

const iscomponent = (item: any) => {
  return item.children && item.children.length > 0
    ? AsideSubMenu
    : AsideItemMenu;
};
</script>

<style lang="scss" scoped>
.aside-menu {
    border-right: none !important;
}
</style>