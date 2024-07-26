import type { Response } from '../types/email'

const fetchEmails = async (term: string, page?: number) => {
    try {
        const url = page !== undefined 
          ? `http://localhost:3000/search/${term}-${page}` 
          : `http://localhost:3000/search/${term}-{}`;
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