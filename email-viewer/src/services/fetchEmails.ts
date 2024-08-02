import type { Response } from '../types/email'
import {SERVER, SEARCH_EMAILS_ENDPOINT} from '../constants/constants'

const fetchEmails = async (term: string, page?: number) => {
    try {
        const url = page !== undefined 
          ? `${SERVER}${SEARCH_EMAILS_ENDPOINT}${term}-${page}` 
          : `${SERVER}${SEARCH_EMAILS_ENDPOINT}${term}-{}`;
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        const data: Response = await response.json();
        if (data) return data?.hits;
    } catch (error) {
        console.error('Error fetching emails', error);
    }
}

export default fetchEmails;