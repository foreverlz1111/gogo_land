function Header() {
    return (<header>
        <div className="ml-4 mr-4 lg:items-center lg:justify-between border-b-4 mb-4">
            <div className="min-w-0 flex-1 mb-4">
                <h2 className="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
                    HintTitle
                </h2>
                <div className="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap sm:space-x-6">
                    <div className="mt-2 flex items-center text-sm text-gray-500">

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