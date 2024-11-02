function Footer() {
    const yes = (<>&#x2705;</>)
    const no = (<>&#x2716;</>)
    const todo = [
        {ready: yes, text: "剥离head,app,foot组件"},
        {ready: no, text: "移动设备分辨率下：点击菜单栏（从左往右弹出）（无回弹）"},
        {ready: no, text: "请求数据接口"},
    ]
    return (
        <>
            <footer className="sm:ml-4 sm:mr-4 text-center bg-white text-white border-t-4 mt-4">
                <div className="container-center px-1 pt-2">
                    <div className="flex justify-center mb-6">
                    </div>
                    <div className="border-b border-teal-100 w-11/12 sm:max-w-5xl mx-auto text-black text-left ">
                        <p className="font-bold mb-4">To-Do:</p>
                        <ul className="list-disc list-inside list-none">
                            {todo.map((e, i) => (<li key={e.text}>{e.ready}{e.text}</li>))}
                        </ul>
                        <div
                            className="grid gap-4 grid-cols-3 text-sm sm:text-lg w-11/12 sm:max-w-5xl mx-auto text-sm sm:text-lg rounded-lg">
                        </div>
                    </div>
                </div>
            </footer>
        </>
    )
}

export default Footer