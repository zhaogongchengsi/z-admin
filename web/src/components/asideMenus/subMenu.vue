<template>
  <el-sub-menu :index="index">
    <template #title>
      <el-icon>
        <component :is="menuInfo.icon" />
      </el-icon>
      <span>{{menuInfo.label}}</span>
    </template>
    <item-menu v-for="(item, index) in menuInfo.children" :key="index" :menuInfo="item" :index="item.path" />
  </el-sub-menu>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ItemMenu from "./itemMenu.vue";
export default defineComponent({
  name: "AsideSubMenu",
  props: {
    index: {
      type: String,
      default: "",
    },
    menuInfo: {
        type: Object,
        default: () => ({}),
    }
  },
  setup(props) {
    
  },
});
</script>