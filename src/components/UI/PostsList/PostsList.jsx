import React, { useEffect, useState } from "react";
import { fetchPosts, deletePost } from "../../../api";

export default function PostsList({ onEdit, externalPosts }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  async function load() {
    setLoading(true);
    setError(null);
    try {
      const data = await fetchPosts();
      setPosts(data.slice(0, 20));
    } catch (err) {
      setError(err.message || "Ошибка при загрузке");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    load();
  }, []);

  // Merge externalPosts (created/updated locally)
  const merged = React.useMemo(() => {
    if (!externalPosts || externalPosts.length === 0) return posts;
    const map = Object.fromEntries(posts.map(p => [p.id, p]));
    // override or add external posts
    externalPosts.forEach(p => { map[p.id] = p; });
    // keep order: external first, then existing not overridden
    const extIds = externalPosts.map(p => p.id).filter(Boolean);
    const remaining = posts.filter(p => !extIds.includes(p.id));
    return [...externalPosts, ...remaining];
  }, [posts, externalPosts]);

  async function handleDelete(id) {
    if (!window.confirm("Удалить пост?")) return;
    try {
      await deletePost(id);
      // обновляем локально
      setPosts((p) => p.filter(x => x.id !== id));
    } catch (err) {
      alert("Ошибка при удалении: " + (err.message || err));
    }
  }

  if (loading) return <div>Loading posts...</div>;
  if (error) return <div style={{ color: "red" }}>{error}</div>;

  return (
    <div>
      <h2>Posts</h2>
      <ul style={{ paddingLeft: 0 }}>
        {merged.map(post => (
          <li key={post.id} style={{ listStyle: "none", marginBottom: 12, borderBottom: "1px solid #eee", paddingBottom: 8 }}>
            <strong>{post.title}</strong>
            <p>{post.body}</p>
            <button onClick={() => onEdit(post)}>Edit</button>{" "}
            <button onClick={() => handleDelete(post.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
