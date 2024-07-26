<template>
    <div class='flex items-center justify-center rounded-full max-w-[600px] bg-[#0d1829] pr-[10px]'>
      <input 
      class="bg-[#0d1829] rounded-l-full text-[16px] pt-[5px] pb-[5px] pl-[20px] pr-[20px] text-[#fff] w-full"
        type="text" 
        v-model="query" 
        placeholder="Search emails..." 
        @keyup.enter="searchFunction"
      />
      <button v-on:click="searchFunction">
        <img class="w-[25px] h-[25px]" :src="searchIcon" alt="search button" />
      </button>
    </div>
  </template>
  
<script setup lang="ts">
import { ref } from 'vue'
import searchIcon from "@/assets/icons/SearchIcon.svg"

const emit = defineEmits(['search'])
const query = ref('')
const searchFunction = () => {
  emit('search', query.value)
}


</script>
  
  <!-- <style scoped>
  .search-bar {
    display: flex;
    justify-content: center;
    margin: 20px;
  }
  input {
    width: 300px;
    padding: 10px;
    margin-right: 10px;
  }
  button {
    padding: 10px 20px;
  }
  </style> -->