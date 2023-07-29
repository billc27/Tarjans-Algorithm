import { useState } from 'react';

function App() {
    const [result, setResult] = useState("");
    const [runtime, setRuntime] = useState("");
    const [filePath, setFilePath] = useState('');

    const handleFileChange = (event) => {
        if (event.target.files.length > 0) {
          setFilePath(event.target.files[0].name);
        }
    };

    function getContent() {
        // Check which input mode is selected
        const inputMode = document.querySelector('input[name="mode"]:checked').value;
    
        if (inputMode === "text") {
            // Get content from text area
            const content = document.querySelector("textarea").value;
            
            // console.log(content);
            return Promise.resolve(content);
        } else {
            // Get content from file input
            const file = document.querySelector('input[type="file"]').files[0];
            return new Promise(resolve => {
                const reader = new FileReader();
                reader.onload = () => {
                    // console.log(reader.result);
                    resolve(reader.result);
                }
                reader.readAsText(file);
            });
        }
    }
    
    function handleVisualizeClick() {
        // Record start time
        const startTime = performance.now();

        getContent().then(content => {
            // Make HTTP request to backend
            fetch("http://localhost:8080/visualize", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ content: content })
            })
            .then(response => response.blob())
            .then(blob => {
                // Create object URL from blob
                const imageUrl = URL.createObjectURL(blob);

                setResult(<img className="mx-auto block max-w-full max-h-full" src={imageUrl} alt="input"/>)

                // Record end time
                const endTime = performance.now();

                // Runtime
                const runtime = endTime - startTime;

                setRuntime(runtime);
            });
        });
    }
    
    function handleFindSccClick() {
        // Record start time
        const startTime = performance.now();

        getContent().then(content => {
            // Make HTTP request to backend
            fetch("http://localhost:8080/scc", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ content: content })
            })
            .then(response => response.blob())
            .then(blob => {
                // Create object URL from blob
                const imageUrl = URL.createObjectURL(blob);
                setResult(<img className="mx-auto block max-w-full max-h-full" src={imageUrl} alt="scc"/>)

                // Record end time
                const endTime = performance.now();

                // Runtime
                const runtime = endTime - startTime;

                setRuntime(runtime);
            });
        });
    }
    
    function handleFindBridgeClick() {
        // Record start time
        const startTime = performance.now();

        getContent().then(content => {
            // Make HTTP request to backend
            fetch("http://localhost:8080/bridge", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ content: content })
            })
            .then(response => response.blob())
            .then(blob => {
                // Create object URL from blob
                const imageUrl = URL.createObjectURL(blob);
                setResult(<img className="mx-auto block max-w-full max-h-full" src={imageUrl} alt="bridge" />)

                // Record end time
                const endTime = performance.now();

                // Runtime
                const runtime = endTime - startTime;

                setRuntime(runtime);
            });
        });
    }
    

    return (
        <div className="App">
            <p className="text-center p-4 text-2xl border bg-title-color border-black mx-auto mt-11 rounded-lg font-serif" style={{maxWidth: "35%"}}>SCC and Bridge Finder Using Tarjan's Algorithm</p>
            <div className="flex mt-12 mb-11">
                <div className="w-1/2 ml-20">

                    {/* Left Section */}
                    <div className="flex items-center">
                        <div className="w-1/2 h-7 mr-2 pl-2 border border-white ml-10 text-wht-color">
                            {filePath}
                        </div>
                        <input
                            className="bg-title-color"
                            type="file"
                            onChange={handleFileChange}
                            style={{ display: 'none' }}
                            id="fileInput"
                        />
                        <button
                            className="border h-7 px-2 bg-title-color rounded-sm hover:bg-gray-200 font-sans"
                            onClick={() => document.getElementById('fileInput').click()}
                        >
                            Browse
                        </button>
                    </div>
                    <div className="flex flex-col items-start w-full pl-10 pr-5">
                        <p className="text-left text-wht-color font-semibold font-sans mt-6">Your Text 
                        </p>
                        <textarea
                            className="w-3/4 p-2 bg-box-color border border-white rounded-sm mt-3 h-80 text-white font-sans"
                            placeholder="Enter text here"
                        />

                        {/* Input Mode */}
                        <div className="mt-2">
                            <p className="inline-block mr-4 text-white font-sans font-semibold">Input: </p>
                            <label className="text-white font-sans font-semibold">
                                <input type="radio" name="mode" value="file" 
                                    className="mr-2 "
                                />
                                File
                            </label>
                            <label className="text-white font-sans font-semibold">
                                <input type="radio" name="mode" value="text" 
                                    className="ml-4 mr-2" 
                                />
                                Text
                            </label>
                        </div>

                        {/* Visualize Button */}
                        <button className="w-3/4 text-center h-9 border bg-box-color border-white mt-4 rounded-sm hover:bg-gray-300 font-sans"
                        onClick={handleVisualizeClick}>
                                Visualize
                        </button>

                        {/* Find SCC Button */}
                        <button className="text-center w-3/4 h-9 border bg-title-color border-white mt-2 rounded-sm hover:bg-gray-300 font-sans"
                        onClick={handleFindSccClick}>
                            Find SCC
                        </button>

                        {/* Find Bridge Button */}
                        <button className="text-center w-3/4 h-9 border bg-title-color border-white mt-2 rounded-sm hover:bg-gray-300 font-sans"
                        onClick={handleFindBridgeClick}>
                            Find Bridge
                        </button>
                    </div>
                    
                </div>

                {/* Right Section */}
                <div className="w-1/2 ml-24 mt-8">
                    <p className="w-4/5 font-sans font-semibold text-center text-wht-color">Result</p>
                    
                    {/* Result Box */}
                    <div className="w-4/5 h-full bg-box-color border border-wht-color rounded-sm mt-3" style={{maxHeight: "87%"}}>
                        {result}
                    </div>

                    <p className="mt-2 font-sans text-wht-color font-semibold">Runtime: {runtime} ms</p>
                </div>
            </div>
        </div>
    );
}

export default App;
