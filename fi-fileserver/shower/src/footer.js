function Footer() {
    return (
        <footer className="ml-4 mr-4 text-center bg-white text-white border-t-4 mt-4">
            <div className="container-center px-1 pt-2">
                <div className="flex justify-center mb-6">
                </div>
                <div className="border-b border-teal-100 w-11/12 sm:max-w-5xl mx-auto text-black text-left ">
                    <p className="font-bold mb-4">To-Do:</p>
                    <ul className="list-disc list-inside">
                        <li></li>
                    </ul>
                    <div
                        className="grid gap-4 grid-cols-3 text-sm sm:text-lg w-11/12 sm:max-w-5xl mx-auto text-sm sm:text-lg rounded-lg">
                    </div>
                </div>
            </div>
        </footer>
    )
}

export default Footer