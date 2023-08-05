import { useCallback, useEffect } from "react";
import ReactFlow, {
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  addEdge,
  Node,
  Connection,
} from "reactflow";
import "reactflow/dist/style.css";
import { useQueryNodes } from "@/hooks/useQueryNodes";

// interface NodeData {
//   label: string;
// }

// interface CustomNode extends Node {
//   data: NodeData;
// }

// const initialNodes: CustomNode[] = [
//   {
//     id: "start",
//     type: "input",
//     position: { x: 0, y: 0 },
//     data: { label: "start" },
//   },
//   {
//     id: "middle",
//     type: "",
//     position: { x: 0, y: 100 },
//     data: { label: "middle" },
//   },
//   {
//     id: "end",
//     type: "output",
//     position: { x: 0, y: 200 },
//     data: { label: "end" },
//   },
// ];

// const initialEdges = [
//   { id: "e1", source: "1", target: "1" },
//   { id: "e2", source: "2", target: "3" },
// ];

function Flow() {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
  const { data: allNodes } = useQueryNodes();

  const onConnect = useCallback(
    (connection: Connection) => setEdges((eds) => addEdge(connection, eds)),
    [setEdges]
  );

  const onNodeDragStop = useCallback((_: MouseEvent, node: Node) => {
    // ドロップされたノードの情報を取得
    console.log("Dropped Node:", node);
  }, []);

  useEffect(() => {
    if (allNodes != undefined) {
      setNodes(allNodes.data["node"]);
      setEdges(allNodes.data["edge"]);
    }
  }, [allNodes]);

  return (
    <div style={{ height: "100vh" }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onNodeDragStop={onNodeDragStop}
        onConnect={onConnect}
        fitView
      >
        <MiniMap />
        <Controls />
        <Background />
      </ReactFlow>
    </div>
  );
}

export default Flow;
