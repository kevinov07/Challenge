<template>
  <div v-if="email" class="email-details">
    <h2>{{ email.subject }}</h2>
    <p><strong>From:</strong> {{ email.from }}</p>
    <p><strong>To:</strong> {{ email.to }}</p>
    <p><strong>Date:</strong> {{ formattedDate }}</p>
    <p>{{ email.body }}</p>
    <!-- <div class="email-body" v-html="email.body"></div> -->
  </div>
  <div v-else class="flex align-center justify-center">
      <p>No email selected</p>
  </div>
</template>
  
<script setup lang="ts">
import { computed, defineProps } from 'vue';
import type { Email } from '../types/email';

const props = defineProps<{
  email: Email | null;
}>();

const formattedDate = computed(() => {
  if (props.email) {
    return new Date(props.email.date).toLocaleString()
  }
  return '';
});

</script>

<!-- <style scoped>
.email-body {
  white-space: pre-wrap;
}
</style> -->
  