import React, { useEffect, useState, useRef } from "react";
import { IoBookmarkOutline, IoChevronBack, IoChevronForward } from "react-icons/io5";
import { TiDeleteOutline } from "react-icons/ti";

function Popup() {
  const [bookmarks, setBookmarks] = useState([]);
  const [error, setError] = useState("");
  const [title, setTitle] = useState("");
  const [category, setCategory] = useState("");
  const [currUrl, setCurrUrl] = useState("Not Found");
  const [selectedCategory, setselectedCategory] = useState("All");
  const scrollRef = useRef(null);
  const [isForm, setIsForm] = useState(false);

  useEffect(() => {
    const fetchBookmarks = async () => {
      try {
        const response = await fetch("http://localhost:8000/bookmarks");
        if (!response.ok) throw new Error("Could not fetch Bookmarks");
        const data = await response.json();
        setBookmarks(data);
      } catch (err) {
        setError("Could not load bookmarks");
      }
    };
    fetchBookmarks();
  }, []);


  const handleDeleteBookmark = async (id) => {
    try {
      const response = await fetch(`http://localhost:8000/delete/bookmarks/${id}`, {
        method: "DELETE",
      });
      if (!response.ok) throw new Error("Failed to delete bookmark");
      setBookmarks((prev) => prev.filter((bm) => bm.id !== id));
    } catch (err) {
      console.error("Error deleting bookmark:", err);
    }
  };

  

  useEffect(() => {
    const getCurrentTab = async () => {
      try {
        const [tab] = await chrome.tabs.query({
          active: true,
          currentWindow: true,
        });
        if (tab && tab.url) {
          setCurrUrl(tab.url);
        }
      } catch (err) {
        console.error("Error getting current tab:", err);
      }
    };
    getCurrentTab();
  }, []);

  const scrollLeft = () => {
    scrollRef.current?.scrollBy({ left: -100, behavior: "smooth" });
  };

  const scrollRight = () => {
    scrollRef.current?.scrollBy({ left: 100, behavior: "smooth" });
  };

  const categories = [
    "All",
    ...new Set(
      bookmarks
        .map((bm) => (bm.category !== "" ? bm.category : null))
        .filter((cat) => cat && cat.trim() !== "")
    ),
  ];

  const filteredBookmarks =
    selectedCategory === "All"
      ? bookmarks
      : bookmarks.filter((bm) => bm.category === selectedCategory);

  const handleCategoryClick = (category) => {
    setselectedCategory(category);
  };

  const handleAddBookmark = async () => {
    try {
      const response = await fetch("http://localhost:8000/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: title,
          category: category,
          url: currUrl,
        }),
      });
      if (!response.ok) {
        throw new Error("Failed to create a new Bookmark");
      }
      const newBookmark = await response.json();
      setBookmarks((prev) => [...prev, newBookmark]);
      setIsForm(false);
      setTitle("");
      setCategory("");
    } catch (err) {
      console.log("Failed to add a bookmark: ", err);
    }
  };

  return (
    <div className="p-3 w-96 bg-gray-900 h-auto border-0 rounded-lg">
      <div className="flex justify-between items-start">
        <h1 className="text-2xl font-bold text-red-500 mb-6 flex items-center">
          <span className="mr-2">🔖</span>
          Dex
        </h1>
        <button
          className="bg-red-500 rounded-lg text-white font-semibold text-md hover:bg-red-400 p-1.5"
          onClick={() => setIsForm(!isForm)}
        >
          Add bookmark
        </button>
      </div>

      {isForm && (
        <div className="fixed inset-0 flex items-center justify-center z-50">
          <div className="bg-gray-800 p-5 rounded-lg w-[90%] max-w-md shadow-xl border border-gray-700">
            <h2 className="text-xl font-bold text-white mb-4">Add Bookmark</h2>
            <form
              onSubmit={(e) => {
                e.preventDefault();
                handleAddBookmark();
              }}
              className="space-y-3"
            >
              <div>
                <label className="block text-sm font-medium text-white mb-1">Title</label>
                <input
                  type="text"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  placeholder="Bookmark Title"
                  className="w-full px-3 py-2 bg-gray-700 text-white rounded-lg focus:ring-2 focus:outline-none focus:ring-red-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-white mb-1">Category</label>
                <input
                  type="text"
                  value={category}
                  onChange={(e) => setCategory(e.target.value)}
                  placeholder="e.g. Work, Personal"
                  className="w-full px-3 py-2 bg-gray-700 text-white rounded-lg focus:ring-2 focus:outline-none focus:ring-red-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-white mb-1">URL</label>
                <input
                  type="text"
                  value={currUrl}
                  disabled
                  className="w-full px-3 py-2 bg-gray-700 text-white rounded-lg opacity-60 cursor-not-allowed"
                />
              </div>

              <div className="flex justify-between mt-4">
                <button
                  type="button"
                  onClick={() => setIsForm(false)}
                  className="px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-500"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-red-500 text-white font-semibold rounded-lg hover:bg-red-400"
                >
                  Create
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {error && (
        <div className="p-3 mb-4 bg-red-100 border border-red-400 text-red-700 rounded-md text-sm">
          {error}
        </div>
      )}

      <div className="mb-4 w-full flex items-center gap-2">
        <button
          onClick={scrollLeft}
          className="text-gray-400 hover:text-white p-1 rounded-full bg-gray-800"
        >
          <IoChevronBack />
        </button>

        <div
          ref={scrollRef}
          className="flex gap-3 overflow-x-hidden scroll-smooth w-full"
        >
          {categories.map((cat, idx) => (
            <button
              key={idx}
              onClick={() => handleCategoryClick(cat)}
              className={`px-2 py-2 whitespace-nowrap rounded-full text-sm font-semibold transition-all duration-150 ${
                selectedCategory === cat
                  ? "bg-red-500 text-white"
                  : "bg-gray-700 text-gray-300 hover:bg-gray-600"
              }`}
            >
              {cat}
            </button>
          ))}
        </div>

        <button
          onClick={scrollRight}
          className="text-gray-400 hover:text-white p-1 rounded-full bg-gray-800"
        >
          <IoChevronForward />
        </button>
      </div>

      <div className="max-h-64 overflow-y-auto pr-1">
        <ul className="space-y-2">
          {filteredBookmarks.map((bm) => (
            <li
              key={bm.id}
              className="group transition-all duration-150 hover:bg-gray-800 rounded-lg p-1"
            >
              <div className="flex items-start gap-2">
                <IoBookmarkOutline className="flex-shrink-0 text-red-500 mt-1" />
                <div className="flex-1">
                  <div className="flex justify-between items-start">
                    <h3 className="font-medium text-gray-100 text-sm font-['Poetsen_One'] truncate">
                      {bm.title}
                    </h3>
                    <div className="hidden group-hover:block">
                      <TiDeleteOutline className="text-orange-400 text-lg cursor-pointer" onClick={() => handleDeleteBookmark(bm.id)} />
                      
                    </div>
                  </div>
                  <a
                    href={bm.url}
                    className="text-blue-400 hover:text-blue-300 text-sm truncate block transition-colors"
                    target="_blank"
                  >
                    {bm.url}
                  </a>
                </div>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default Popup;
