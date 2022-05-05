<template>
  <el-menu>
      <component 
        v-for="(item, index) in porps.menuOptions" 
        :key="index" 
        :is="iscomponent(item)" 
        :menuInfo="item"
       ></component>
  </el-menu>
</template>

<script lang="ts" setup>
import AsideItemMenu from "./itemMenu.vue"
import AsideSubMenu from "./subMenu.vue"

const porps = defineProps({
  menuOptions: {
    type: Array,
    default: () => [],
  },
});

const iscomponent = (item: any) => {
    return item.children && item.children.length > 0 ? AsideSubMenu : AsideItemMenu;
}


</script>