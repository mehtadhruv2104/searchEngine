import { useState } from "react";
import { searchQueryInterface } from "./interface";


const SearchBox = (props:{onSearch:Function, setSearchQuery:Function, searchQuery:searchQueryInterface}) => {


    const {searchQuery, setSearchQuery, onSearch} = props;

    const handleChange = (e) => {
        const { name, value } = e.target;
        setSearchQuery(prev => ({ ...prev, [name]: value }));
    };

    const handleSearch =() =>{
        onSearch && onSearch();
    }

    const handleClearAll = () =>{
        setSearchQuery({
            query:"",
            severity:"",
            msg_id:"",
            namespace:"",
            app_name:"",
        })
    }

    return(
        <div className="flex flex-col space-y-6 w-full max-w-5xl mx-auto p-4">
            <div className="w-full">
                <input
                    type="text"
                    name="query"
                    value={searchQuery.query || ''}
                    onChange={handleChange}
                    placeholder="Enter search query..."
                    className="w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
            </div>
            <div className="grid grid-cols-5 gap-4">

                <input
                type="text"
                name="severity"
                placeholder="Severity"
                value={searchQuery.severity || ''}
                onChange={handleChange}
                className="px-4 py-2 border rounded shadow"
                />
                <input
                type="text"
                name="app_name"
                placeholder="App Name"
                value={searchQuery.app_name || ''}
                onChange={handleChange}
                className="px-4 py-2 border rounded shadow"
                />
                <input
                type="text"
                name="namespace"
                placeholder="Namespace"
                value={searchQuery.namespace || ''}
                onChange={handleChange}
                className="px-4 py-2 border rounded shadow"
                />
                <input
                type="text"
                name="msg_id"
                placeholder="Msg ID"
                value={searchQuery.msg_id || ''}
                onChange={handleChange}
                className="px-4 py-2 border rounded shadow"
                />
            </div>
            <div className="flex justify-center flex-row gap-x-8">
                <button
                onClick={handleSearch}
                className="px-6 py-3 bg-blue-600 text-white font-medium cursor-pointer rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 shadow-md transition duration-200"
                >
                Search
                </button>
                <button
                onClick={handleClearAll}
                className="px-6 py-3 bg-blue-600 text-white font-medium cursor-pointer rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 shadow-md transition duration-200"
                >
                Clear all 
                </button>
            </div>
        </div> 
    )
}

export default SearchBox;