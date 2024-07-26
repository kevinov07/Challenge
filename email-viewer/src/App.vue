<template>
  <div class="flex flex-col gap-[20px] px-[30px] py-[30px]">
    <header class="w-full flex items-center justify-start">
      <h1 class="flex gap-[5px] items-center text-2xl pl-12">
        <img class="h-[auto] max-h-[30px]" :src="EmailIcon" alt="Email icon"/>
        <span>Search Email</span>
      </h1>
    </header>
    <SearchBar class="w-full max-w-3xl mx-auto"  @search="handleSearch"/>
    <div v-if="emails.length" class="flex w-full">
      <div class="flex-1">
        <EmailTable :emails="emails" @selectEmail="selectEmail"/>
      </div>
      <div class="flex-1 ml-[20px]">
        <EmailDetails :email="selectedEmail || null"/>
      </div>
    </div>
    <div v-else class="text-[#A0A0A0] flex align-center justify-center">
      No emails found.
    </div>
  </div>
</template>

<script setup lang="ts"> 
import { ref } from 'vue';
import SearchBar from './components/SearchBar.vue'
import EmailTable from './components/EmailTable.vue'
import EmailDetails from './components/EmailDetails.vue'
import EmailIcon from '@/assets/icons/EmailIcon.svg'
import fetchEmails from './services/fetchEmails'; './services/EmailService';
import type { Email, Hit } from './types/email';

const emails = ref<Email[]>([]);
const selectedEmail = ref<Email>();


const handleSearch = async (query:string) => {
  try {
    const results = await fetchEmails(query);
    emails.value = results?.hits.map(hit => hit._source) ?? [];
  } catch (error) {
    console.error('Error fetching emails:',error);
  }
}

const selectEmail = (email: Email) => {
    selectedEmail.value = email;
}
</script>