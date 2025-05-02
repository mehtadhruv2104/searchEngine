import { searchQueryInterface } from './interface'
import { searchLogs } from '../service/searchService'
//import "./search.css"
import Header from './Header'
import { useActionState, useState } from 'react'
import SearchBox from './SearchBox'
import Loading from './Loading'
import Results from './Results'



const SearchEngine = () =>{

    const [searchQuery, setSearchQuery] = useState<searchQueryInterface>({});
    console.log("search query", searchQuery);
    const [searchTime, setSearchTime] = useState(0);
    const [totalMatches, setTotalMatches] = useState(0);
    const [results, setResults] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [showResults, setShowResults] = useState<boolean>(false)

    const handleSearch = async () => {
        setLoading(true);
        setError(null);
        try{
            const data = await searchLogs(searchQuery);
            setResults(data?.records || []);
            setSearchTime(data?.time_ms || 0);
            setTotalMatches(data?.count || 0);
        }catch (err) {
            setError(err.message);
          } finally {
            setLoading(false);
            setShowResults(true)
        }
    }



    
    return(
        <div className='searchEngine'>
            <div className='flex flex-col gap-4'>
                <Header />
                <button onClick={handleSearch}>Click</button>
                <SearchBox onSearch={handleSearch} setSearchQuery={setSearchQuery} searchQuery={searchQuery}/>
                {loading && <Loading/>}
                {showResults && (
                    <Results results={results} searchTime={searchTime} totalMatches={totalMatches}/>
                )}
            </div>
        </div>
    )
}

export default SearchEngine;