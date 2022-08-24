<script lang="ts" setup>

import {onMounted, reactive, ref} from "@vue/runtime-core";

const word = ref<string>("コンテンツ")
onMounted(() => word.value = word.value.toUpperCase())

// const style = ref('w-3/4');
interface Style {
  css     : string;
  isSelect: boolean;
}
const style = reactive<Style>({
  css     : 'w-5/6',
  isSelect: false
})
const styleChange = (target: boolean) => {
  style.isSelect = target
  style.css      = target ? 'w-4/6' : 'w-5/6'
}

</script>

<template>
  <div>
    <DashboardHeader
      @open-hamburger-menu="styleChange"
    />
    <div class="flex flex-row w-full">
      <DashboardSidebar
        class="w-1/6"
      />
      <div class="min-h-screen flex items-center justify-center " v-bind:class="style.css">
        <div class="flex flex-col">
          <div>{{word}}</div>
          <div class="text-center text-white duration-700 transform hover:scale-125 transition w-max bg-indigo-600">送信</div>
        </div>

      </div>
      <div class="w-1/6 flex flex-col items-center justify-start bg-pink-600" v-if="style.isSelect">
        TEST
      </div>
    </div>
  </div>
</template>


