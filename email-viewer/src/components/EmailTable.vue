<template>
<div class="px-4 py-3 @container">
  <div v-if="emails && emails.length" class="flex rounded-xl border border-[#D5D6DD] bg-[#F9FAFA]">
    <table class="min-w-full max-w-[800px] table-fixed">
      <thead>
        <tr class="bg-[#FFFFFF]">
          <th class=" px-4 py-3 text-left text-[#1C1D22] w-[160px] text-sm font-medium leading-normal">
            Subject
          </th>
          <th class="px-4 py-3 text-left text-[#1C1D22] w-[160px] text-sm font-medium leading-normal">
          From
          </th>
          <th class="px-4 py-3 text-left text-[#1C1D22] w-[160px] text-sm font-medium leading-normal">
            To
          </th>
          <th class="px-4 py-3 text-left text-[#1C1D22] w-[160px] text-sm font-medium leading-normal">
            Date
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(email, index) in emails" :key="index" @click="selectEmail(email)" class="border-t border-t-[#D5D6DD] cursor-pointer hover:bg-[#EAEAEA]">
          <td class="h-[72px] px-4 py-2 w-[400px] text-[#1C1D22] text-sm font-normal leading-normal truncate">
            {{ email.subject }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[400px] text-[#1C1D22] text-sm font-normal leading-normal truncate">
            {{ email.from }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[400px] text-[#1C1D22] text-sm font-normal leading-normal truncate">
            {{ email.to }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[400px] text-[#3C3F4A] text-sm font-normal leading-normal truncate">
            {{ new Date(email.date).toLocaleString() }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[400px] text-[#3C3F4A] text-sm font-normal leading-normal truncate">
            {{ email.subject }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  <div v-else class="text-center text-gray-500">
    No emails found.
  </div>
</div>
</template>
  
<script setup lang="ts">
import { defineProps, defineEmits, computed } from 'vue'
import type { Email, Hit } from '@/types/email';

const props = defineProps<{emails: Email[]}>()
const emit = defineEmits<{
  (e: 'selectEmail', email: Email): void;
}>();

const selectEmail = (email: Email) => {
  emit('selectEmail', email)
};

  
</script>
  
<style scoped>
/* .email-table {
  flex: 1;
  margin: 20px;
  overflow-x: auto;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  padding: 10px;
  border: 1px solid #ddd;
  text-align: left;
}
th {
  background-color: #f4f4f4;
}
tr:hover {
  background-color: #f1f1f1;
  cursor: pointer;
} */
</style>
  