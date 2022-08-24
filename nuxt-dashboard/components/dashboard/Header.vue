<script lang="ts" setup>
import {ref} from "@vue/runtime-core";

interface Emits {
  (e: 'open-hamburger-menu', target: boolean): void
}
const emit = defineEmits<Emits>();

const isOpen         = ref(false);
const hamburgerClick = () => {
  isOpen.value = !isOpen.value
  emit('open-hamburger-menu', isOpen.value)
}
</script>

<template>
  <div class="w-full flex flex-row bg-green-700 py-3">
    <div class="flex justify-between items-center w-full px-8 text-white">
      <div>左</div>
      <div class="cursor-pointer p-1" v-on:click="hamburgerClick()">
        <svg class="h-6 w-6 fill-current" viewBox="0 0 24 24">
          <path d="M24 6h-24v-4h24v4zm0 4h-24v4h24v-4zm0 8h-24v4h24v-4z"/>
        </svg>
      </div>
    </div>
  </div>
</template>
