import { useCallback } from "react";
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

interface NodeData {
  label: string;
}

interface CustomNode extends Node {
  data: NodeData;
}

const initialNodes: CustomNode[] = [
  {
    id: "start",
    type: "input",
    position: { x: 0, y: 0 },
    data: { label: "start" },
  },
  {
    id: "middle",
    position: { x: 0, y: 100 },
    data: { label: "middle" },
  },
  {
    id: "end",
    type: "output",
    position: { x: 0, y: 200 },
    data: { label: "end" },
  },
];

const initialEdges = [
  { id: "e1", source: "start", target: "middle" },
  { id: "e2", source: "middle", target: "end" },
];

function Flow() {
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);

  const onConnect = useCallback(
    (connection: Connection) => setEdges((eds) => addEdge(connection, eds)),
    [setEdges]
  );

  const onNodeDragStop = useCallback((event: MouseEvent, node: Node) => {
    // ドロップされたノードの情報を取得
    console.log("Dropped MouseEvent:", event);
    console.log("Dropped Node:", node);
  }, []);

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
