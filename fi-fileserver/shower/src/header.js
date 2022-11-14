function Header() {
    return (<header>
        <div className="ml-4 mr-4 lg:items-center lg:justify-between border-b-4 mb-4">
            <div className="min-w-0 flex-1 mb-4">
                <h2 className="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
                    HintTitle
                </h2>
                <div className="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap sm:space-x-6">
                    <div className="mt-2 flex items-center text-sm text-gray-500">

                        <svg aria-hidden="true" className="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400"
                             fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                            <path clip-rule="evenodd"
                                  d="M5.75 2a.75.75 0 01.75.75V4h7V2.75a.75.75 0 011.5 0V4h.25A2.75 2.75 0 0118 6.75v8.5A2.75 2.75 0 0115.25 18H4.75A2.75 2.75 0 012 15.25v-8.5A2.75 2.75 0 014.75 4H5V2.75A.75.75 0 015.75 2zm-1 5.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h10.5c.69 0 1.25-.56 1.25-1.25v-6.5c0-.69-.56-1.25-1.25-1.25H4.75z"
                                  fill-rule="evenodd"/>
                        </svg>
                        Date
                    </div>
                </div>
            </div>
            <div className="mt-5 flex lg:mt-0 lg:ml-4 mb-4">
            <span className="hidden sm:block">
                <button
                    className="inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-300 shadow-sm bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                    disabled
                    type="button">
                    Edit
      </button>
    </span>
                <span className="sm:ml-3">
      <button
          className="inline-flex items-center rounded-md border border-transparent px-4 py-2 text-sm font-medium bg-gray-300 text-white shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          disabled
          type="button">
          Upload
      </button>
    </span>
                <div className="relative ml-3 sm:hidden">
                </div>
            </div>
        </div>
    </header>)
}

export default Header