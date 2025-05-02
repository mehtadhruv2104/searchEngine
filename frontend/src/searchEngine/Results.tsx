
import React from 'react';

const Results = ({ results, searchTime, totalMatches }) => {

  if (!results || results.length === 0) {
    return (
      <div
        className="text-center py-8 text-gray-500"
        aria-live="polite"
      >
        No results found.
      </div>
    );
  }

  return (
    <div className="py-6" aria-live="polite">

      <div className="mb-6 bg-blue-50 border border-blue-200 rounded-lg p-4 shadow-sm">
        <p className="text-sm font-medium text-blue-800">
          Found <span className="font-bold">{totalMatches}</span> matches in{' '}
          <span className="font-bold">{searchTime}</span>ms
        </p>
      </div>

      <ul className="space-y-4">
        {results.map((result, index) => (
          <li
            key={index}
            className="p-4 bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200"
          >
            
            <p className="text-sm text-gray-600 mt-1">
              <span className="font-medium">Sender:</span> {result?.Sender}
            </p>
            <p className="text-sm text-gray-600">
              <span className="font-medium">Timestamp:</span>{' '}
              {new Date(result?.NanoTimeStamp / 1000000).toLocaleString()}
            </p>
            <p className="text-sm text-gray-600">
              <span className="font-medium">Event:</span> {result?.Event}
            </p>
            <p className="text-sm text-gray-600">
              <span className="font-medium">AppName:</span> {result?.AppName}
            </p>
            <p className="text-sm text-gray-600">
              <span className="font-medium">Message:</span> {result?.Message}
            </p>
            
            
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Results;





