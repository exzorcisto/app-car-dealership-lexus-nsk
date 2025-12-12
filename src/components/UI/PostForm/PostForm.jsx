import React, { useState } from "react";
import { createPost } from "../../../api";

export default function PostForm({ onAdded }) {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e) {
    e.preventDefault();
    if (!title.trim() || !body.trim()) {
      alert("Заполните заголовок и тело");
      return;
    }
    setLoading(true);
    try {
      const newPost = await createPost({ title, body, userId: 1 });
      onAdded(newPost);
      setTitle("");
      setBody("");
    } catch (err) {
      alert("Ошибка при создании: " + (err.message || err));
    } finally {
      setLoading(false);
    }
  }

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: 24 }}>
      <h3>Create post</h3>
      <div>
        <input placeholder="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
      </div>
      <div>
        <textarea placeholder="Body" value={body} onChange={(e) => setBody(e.target.value)} />
      </div>
      <button type="submit" disabled={loading}>
        {loading ? "Saving..." : "Add Post"}
      </button>
    </form>
  );
}
