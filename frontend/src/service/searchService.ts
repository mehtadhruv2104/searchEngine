import { searchQueryInterface } from "../searchEngine/interface";


export async function searchLogs(data:searchQueryInterface) {
   try{
    const response = await fetch(`${import.meta.env.VITE_BASE_URL}/api/search`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          query: data?.query,
          severity: data?.severity,
          app_name: data?.app_name,
          namespace: data?.namespace,
          msg_id: data?.msg_id,
        }),
      });
    
      if (!response.ok) {
        throw new Error('Internal Server Error');
      }
    
      return await response.json();
   }catch(error){
    console.error('Search error:', error);
    return null;
   }
  }

