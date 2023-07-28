import { useState } from 'react';



function App() {
    const [filePath, setFilePath] = useState('');

    const handleFileChange = (event) => {
        if (event.target.files.length > 0) {
          setFilePath(event.target.files[0].name);
        }
    };

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

                        {/* Visualize Button */}
                        <button className="text-center w-3/4 h-9 border bg-box-color border-white mt-4 rounded-sm hover:bg-gray-300 font-sans" >
                            Visualize
                        </button>

                        {/* Find Button */}
                        <button className="text-center w-3/4 h-9 border bg-title-color border-white mt-2 rounded-sm hover:bg-gray-300 font-sans">
                            Find SCC
                        </button>

                        <button className="text-center w-3/4 h-9 border bg-title-color border-white mt-2 rounded-sm hover:bg-gray-300 font-sans">
                            Find Bridge
                        </button>
                    </div>
                    
                </div>

                {/* Right Section */}
                <div className="w-1/2 ml-24 mt-8">
                    <p className="w-4/5 font-sans font-semibold text-center text-wht-color">Result</p>
                    <div class="w-4/5 h-full bg-box-color border border-wht-color rounded-sm mt-3" style={{maxHeight: "87%"}}></div>
                    <p className="mt-2 font-sans text-white size-sm">Runtime: </p>
                </div>
            </div>
        </div>
    );
}

export default App;
