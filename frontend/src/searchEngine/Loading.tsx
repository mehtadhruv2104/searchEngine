import React from 'react';

const Loading = () => {
  return (
    <div
      className="flex justify-center items-center py-8"
      aria-label="Loading search results"
      role="status"
    >
      <div className="h-10 w-10 border-4 border-blue-200 border-t-blue-600 rounded-full animate-spin"></div>
    </div>
  );
};

export default Loading;