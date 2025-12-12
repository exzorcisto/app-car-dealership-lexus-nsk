import React, { useState } from "react";
import { updatePost } from "../../../api";

export default function EditPostForm({ post, onCancel, onSaved }) {
  const [title, setTitle] = useState(post.title || "");
  const [body, setBody] = useState(post.body || "");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e) {
    e.preventDefault();
    setLoading(true);
    try {
      const updated = await updatePost(post.id, { ...post, title, body });
      onSaved(updated);
    } catch (err) {
      alert("Ошибка при обновлении: " + (err.message || err));
    } finally {
      setLoading(false);
    }
  }

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: 16 }}>
      <h3>Edit post #{post.id}</h3>
      <div>
        <input value={title} onChange={(e) => setTitle(e.target.value)} />
      </div>
      <div>
        <textarea value={body} onChange={(e) => setBody(e.target.value)} />
      </div>
      <button type="submit" disabled={loading}>
        {loading ? "Saving..." : "Save"}
      </button>{" "}
      <button type="button" onClick={onCancel}>Cancel</button>
    </form>
  );
}
