<template>
  <div class="flex flex-col gap-[20px] align-center justify-center">
    <header class="w-[fit-content]">
      <h1 class="flex gap-[5px] align-center justify-center ml-[20px] mt-[20px]">
        <img class="h-[auto] max-h-[30px] mt-[5px]" :src="EmailIcon" alt="Email icon"/>
        <span class="text-[#16b648]">Search Email</span>
      </h1>
    </header>
    <SearchBar @search="handleSearch"/>
    <div  class="flex w-full">
      <div class="flex-1">
        <EmailTable :emails="emails" @selectEmail="selectEmail"/>
      </div>
      <div class="flex-1 ml-[20px]">
        <EmailDetails :email="selectedEmail || null"/>
      </div>
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