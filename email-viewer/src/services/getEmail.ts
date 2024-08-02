import type { Response } from '../types/email';
import { EMAIL_ENDPOINT, SERVER } from '../constants/constants';


const getEmail = async (id: string) => {
    try {
        const response = await fetch(`${SERVER}${EMAIL_ENDPOINT}${id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        const data: Response = await response.json();
        if (data.hits.hits.length === 1) return data?.hits.hits[0]._source;
    } catch (error) {
        console.error('Error fetching email', error);
    }
}




export default getEmail;

    