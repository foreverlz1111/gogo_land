function Header() {
    return (<header>
        <div className="sm:ml-4 sm:mr-4 lg:items-center lg:justify-between border-b-4 mb-4">
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

        </div>
    </header>)
}

export default Header