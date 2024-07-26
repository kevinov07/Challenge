import type { Response } from '../types/email';


const getEmail = async (id: string) => {
    try {
        const response = await fetch(`http://localhost:3000/email/${id}`, {
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

    