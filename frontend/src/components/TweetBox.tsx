
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Camera, Image, MapPin, Smile } from "lucide-react";

interface TweetBoxProps {
  onTweetSubmit: (content: string) => Promise<void>;
  isSubmitting?: boolean;
}

export function TweetBox({ onTweetSubmit, isSubmitting = false }: TweetBoxProps) {
  const [tweetContent, setTweetContent] = useState("");
  
  const handleSubmit = async () => {
    if (!tweetContent.trim()) return;
    
    await onTweetSubmit(tweetContent);
    setTweetContent("");
  };

  return (
    <div className="tweet-box bg-card">
      <div className="mb-2">
        <Textarea
          placeholder="O que está acontecendo?"
          className="min-h-[120px] border-none focus:ring-1 focus:ring-go-blue bg-transparent resize-none"
          value={tweetContent}
          onChange={(e) => setTweetContent(e.target.value)}
        />
      </div>
      
      <div className="flex items-center justify-between border-t pt-3 mt-2">
        <div className="flex space-x-1">
          <button className="go-icon-button" title="Imagem">
            <Image className="h-5 w-5" />
          </button>
          <button className="go-icon-button" title="GIF">
            <span className="font-semibold text-sm">GIF</span>
          </button>
          <button className="go-icon-button" title="Enquete">
            <span className="font-semibold text-sm">POLL</span>
          </button>
          <button className="go-icon-button" title="Emoji">
            <Smile className="h-5 w-5" />
          </button>
          <button className="go-icon-button" title="Câmera">
            <Camera className="h-5 w-5" />
          </button>
          <button className="go-icon-button" title="Localização">
            <MapPin className="h-5 w-5" />
          </button>
        </div>
        
        <Button 
          onClick={handleSubmit}
          disabled={!tweetContent.trim() || isSubmitting}
          className="go-button"
        >
          {isSubmitting ? "Postando..." : "Postar"}
        </Button>
      </div>
    </div>
  );
}
