<template>
<div class="px-4 py-3 @container">
  <div class="flex overflow-hidden rounded-xl border border-[#D5D6DD] bg-[#F9FAFA]">
    <table class="w-full table-fixed">
      <thead>
        <tr class="bg-[#FFFFFF]">
          <th class=" px-4 py-3 text-left text-[#1C1D22] w-[200px] text-sm font-medium leading-normal">
            Subject
          </th>
          <th class="px-4 py-3 text-left text-[#1C1D22] w-[200px] text-sm font-medium leading-normal">
          From
          </th>
          <th class="px-4 py-3 text-left text-[#1C1D22] w-[200px] text-sm font-medium leading-normal">
            To
          </th>
          <th class="px-4 py-3 text-left text-[#1C1D22] w-[200px] text-sm font-medium leading-normal">
            Date
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(email, index) in emails" :key="index" @click="selectEmail(email)" class="border-t border-t-[#D5D6DD] cursor-pointer hover:bg-[#EAEAEA]">
          <td class="h-[72px] px-4 py-2 w-[200px] text-[#1C1D22] text-sm font-normal leading-normal overflow-hidden whitespace-nowrap text-ellipsis">
            {{ email.subject }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[200px] text-[#1C1D22] text-sm font-normal leading-normal overflow-hidden whitespace-nowrap text-ellipsis">
            {{ email.from }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[200px] text-[#1C1D22] text-sm font-normal leading-normal overflow-hidden whitespace-nowrap text-ellipsis">
            {{ email.to }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[200px] text-[#3C3F4A] text-sm font-normal leading-normal overflow-hidden whitespace-nowrap text-ellipsis">
            {{ new Date(email.date).toLocaleString() }}
          </td>
          <td class="h-[72px] px-4 py-2 w-[200px] text-[#3C3F4A] text-sm font-normal leading-normal overflow-hidden whitespace-nowrap text-ellipsis">
            {{ email.subject }}
          </td>
        </tr>
      </tbody>
    </table>
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