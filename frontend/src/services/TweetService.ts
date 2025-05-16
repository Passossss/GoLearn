
import { useToast } from "@/components/ui/use-toast";

export interface Tweet {
  id: string;
  description: string;
  created_at?: string;
}

// Mock data for tweets
const mockTweets: Tweet[] = [
  {
    id: "1",
    description: "Bem-vindo ao GoTweet! Esta é uma demonstração do aplicativo.",
    created_at: new Date().toISOString()
  },
  {
    id: "2",
    description: "O GoTweet é inspirado no design do Twitter com a cor do Go (#00ADD8).",
    created_at: new Date(Date.now() - 3600000).toISOString() // 1 hour ago
  },
  {
    id: "3",
    description: "Você pode postar, visualizar e excluir tweets nesta demonstração.",
    created_at: new Date(Date.now() - 7200000).toISOString() // 2 hours ago
  }
];

// In-memory store for the mock implementation
let tweets = [...mockTweets];
let nextId = 4;

export const useTweetService = () => {
  const { toast } = useToast();

  const getTweets = async (): Promise<Tweet[]> => {
    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 500));
    
    try {
      // Return the mock tweets
      return [...tweets];
    } catch (error) {
      console.error("Error fetching tweets:", error);
      toast({
        title: "Error",
        description: "Failed to load tweets. Please try again.",
        variant: "destructive"
      });
      return [];
    }
  };

  const postTweet = async (description: string): Promise<Tweet | null> => {
    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 500));

    try {
      const newTweet: Tweet = {
        id: String(nextId++),
        description,
        created_at: new Date().toISOString()
      };
      
      // Add tweet to the beginning of the array
      tweets = [newTweet, ...tweets];

      toast({
        title: "Success",
        description: "Your tweet was posted successfully!",
      });
      
      return newTweet;
    } catch (error) {
      console.error("Error posting tweet:", error);
      toast({
        title: "Error",
        description: "Failed to post your tweet. Please try again.",
        variant: "destructive"
      });
      return null;
    }
  };

  const deleteTweet = async (id: string): Promise<boolean> => {
    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 500));

    try {
      const initialLength = tweets.length;
      tweets = tweets.filter(tweet => tweet.id !== id);
      
      const deleted = tweets.length < initialLength;
      
      if (deleted) {
        toast({
          title: "Success",
          description: "Tweet deleted successfully!",
        });
      } else {
        toast({
          title: "Warning",
          description: "Tweet not found.",
        });
      }
      
      return deleted;
    } catch (error) {
      console.error("Error deleting tweet:", error);
      toast({
        title: "Error",
        description: "Failed to delete the tweet. Please try again.",
        variant: "destructive"
      });
      return false;
    }
  };

  return {
    getTweets,
    postTweet,
    deleteTweet
  };
};
