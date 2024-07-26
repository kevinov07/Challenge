<template>
  <div class="px-4 py-3">
    <div v-if="email" class="email-details">
      <h2 class="text-3xl font-bold">{{ email.subject }}</h2>
      <p>
        <strong class="text-base">From: </strong> 
        <span class="text-base text-[#b6bbe0] ">{{ email.from }}</span>
      </p>
      <p>
        <strong>To: </strong> 
        <span class="text-base text-[#b6bbe0]">{{ email.to }}</span>
      </p>
      <p>
        <strong>Date: </strong> 
        <span class="text-base text-[#b6bbe0]">{{ new Date(email.date).toLocaleString() }}</span>
      </p>
      <br>
      <p>
        <span class="text-base text-[#b6bbe0]">{{ email.body }}</span>
      </p>
      <!-- <div class="email-body" v-html="email.body"></div> -->
    </div>
    <div v-else class="flex align-center justify-center">
        <p class="text-[#A0A0A0]">No email selected</p>
    </div>
  </div>
</template>
  
<script setup lang="ts">
import { computed, defineProps } from 'vue';
import type { Email } from '../types/email';

const props = defineProps<{
  email: Email | null;
}>();

</script>